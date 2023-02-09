<script lang="ts">
    import { page } from '$app/stores';
    import api from '$lib/api';
    import type { Chip } from '$lib/class/chip';
    import Button from '$lib/components/button.svelte';
    import Chips from '$lib/components/chips.svelte';
    import Content from '$lib/components/content.svelte';
    import { popup } from '$lib/components/modal.svelte';
    import notification from '$lib/components/notification/notification';
    import Textfield from '$lib/components/textfield.svelte';
    import Titlebar from '$lib/components/titlebar.svelte';
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';

    let parameter = '',
        property = 'data',
        query = '';

    let apiID = '';
    let queries: any[] = [];
    let groups: Chip[] = [];

    let method = '';
    let path = '';

    onMount(() => {
        const searchParams = get(page).url.searchParams;
        let val = searchParams.get('path');
        path = val == null ? '' : val;
        val = searchParams.get('method');
        method = val == null ? '' : val;

        api.get(
            '/admin/apis/details?data=queries,groups&path=' +
                path +
                '&method=' +
                method
        ).then((resp) => {
            if (resp.status == 0) {
                if (resp.data.queries != null) {
                    for (const q of resp.data.queries) {
                        const split = <string[]>q.params.split(',');
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
                            { query: query, property: q.property, data: q },
                        ];
                    }
                }

                if (resp.data.groups != null) {
                    groups = [...resp.data.groups, { label: '+' }];
                }
            }
        });
    });

    function addQuery() {
        api.post('/admin/apis/query/add?api_id=' + apiID, {
            query: query,
            params: parameter,
            property: property,
        })
            .then((resp) => {
                if (resp.status == 0) {
                    parameter = '';
                    property = 'data';
                    query = '';
                } else notification.show('Error', resp.message);
            })
            .catch((e) => notification.show('Error', e));
    }

    function clickChip(ev: any) {
        const chip = <Chip>ev.detail;
        switch (chip.data) {
            case undefined:
            case null:
                popup('/admin/apis/add_group', {
                    api_id: apiID,
                });
                break;
        }
    }
</script>

<Titlebar>API :: {method} {path}</Titlebar>
<Content>
    <div class="flex flex-row space-x-2 items-center">
        <div>Allowed Groups:</div>
        <div><Chips data={groups} on:click={clickChip} /></div>
    </div>

    <div class="queries">
        {#each queries as query}
            <div
                class="flex flex-row"
                on:click={() => popup('/admin/page/query', query.data)}
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
            <Button on:click={addQuery}>Save</Button>
        </div>
    </div>
</Content>

<style lang="scss">
    @import '../../../../styles/colors';

    .add {
        @apply p-4 rounded shadow;
        border: 1px solid $color-border;
    }

    .queries {
        @apply rounded cursor-pointer overflow-clip;
        border: 1px solid $color-border;

        &:hover {
            border: 1px solid $color-secondary;
            background-color: transparentize($color-secondary, 0.9);
        }
        & > div {
            & > div {
                @apply px-2 leading-6;
                &:first-child {
                    background-color: $color-border;
                }
            }
        }

        & :global(.param) {
            @apply rounded p-1;
            background-color: $color-secondary;
            color: $color-text;
        }
    }
</style>
