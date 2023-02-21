<script lang="ts">
    import { page } from '$app/stores';
    import api from '$lib/api';
    import type { Chip } from '$lib/class/chip';
    import Button from '$lib/components/button.svelte';
    import Content from '$lib/components/content.svelte';
    import notification from '$lib/components/notification/notification';
    import Textfield from '$lib/components/textfield.svelte';
    import Titlebar from '$lib/components/titlebar.svelte';
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';
    import QueryDialog from './query-dialog.svelte';

    //new query fields
    let parameter = '',
        property = 'data',
        query = '';

    let disabled = false;

    let queries: any[] = [];
    let groups: Chip[] = [];

    let method = '';
    let path = '';
    let queryDialog: QueryDialog;

    onMount(() => {
        const searchParams = get(page).url.searchParams;
        let val = searchParams.get('path');
        path = val == null ? '' : val;
        val = searchParams.get('method');
        method = val == null ? '' : val;
        refresh('queries,groups');
    });

    function refresh(param: string) {
        const url =
            '/api/admin/apis/details?data=' +
            param +
            '&path=' +
            path +
            '&method=' +
            method;
        api.get(url).then((resp) => {
            if (resp.status == 0) {
                if (resp.data.queries != null) {
                    queries = [];
                    for (const q of resp.data.queries) {
                        const split = <string[]>q.parameter.split(',');
                        let query = q.query;

                        for (let s of split) {
                            s = s.trim();
                            query = query.replace(
                                '?',
                                '<span class="param">' + s + '</span>'
                            );
                        }
                        queries = [
                            ...queries,
                            {
                                query: query,
                                property: q.property,
                                data: q,
                            },
                        ];
                    }
                }

                if (resp.data.groups != null) {
                    groups = [];
                    groups = [...resp.data.groups, { label: '+' }];
                }
            }
        });
    }

    function save() {
        disabled = true;
        if (query.trim() == '') return;
        api.postData('/admin/apis/query/add', {
            path: path,
            method: method,
            query: query,
            parameter: parameter,
            property: property,
        })
            .then((data) => {
                refresh('queries');
                notification.show('Success', 'Query saved');
            })
            .catch((e) => notification.error(e))
            .finally(() => (disabled = false));
    }
</script>

<Titlebar>API :: {method} {path}</Titlebar>
<Content>
    <div class="queries">
        {#each queries as query}
            <div
                class="flex flex-row"
                on:click={() => queryDialog.show(query.data)}
                on:keypress={() => {}}
            >
                <div>{query.property}</div>
                <div class="query">{@html query.query}</div>
            </div>
        {/each}
    </div>
    <div class="add">
        <div>Add Query:</div>
        <div class="flex flex-row space-x-2">
            <Textfield label="Parameter" bind:value={parameter} />
            <Textfield label="Property" bind:value={property} />
        </div>
        <Textfield type="textarea" label="Query" bind:value={query} />
        <div class="flex">
            <div class="flex-1" />
            <Button on:click={save}>Save</Button>
        </div>
    </div>
</Content>

<QueryDialog bind:this={queryDialog} on:success={() => refresh('queries')} />

<style lang="scss">
    @import '../../../../styles/colors';

    .add {
        @apply p-4 rounded shadow;
        border: 1px solid $color-border;
    }

    .queries {
        @apply rounded cursor-pointer overflow-clip;
        border: 1px solid $color-border;

        & > div {
            border: 2px solid transparent;
            &:not(:last-of-type) {
                border-bottom: 2px solid $color-border;
            }
            &:hover {
                border-color: $color-secondary;
                background-color: transparentize($color-secondary, 0.9);
            }
            & > div {
                @apply px-2 leading-6;
                &:first-child {
                    background-color: transparentize($color-border, 0.5);
                }
            }
        }

        & :global(.param) {
            @apply rounded px-1 py-0.5;
            background-color: $color-secondary;
            color: $color-text;
        }
    }
</style>
