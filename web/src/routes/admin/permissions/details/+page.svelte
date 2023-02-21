<script lang="ts">
    import { goto } from '$app/navigation';
    import { page } from '$app/stores';
    import api from '$lib/api';
    import type { Chip } from '$lib/class/chip';
    import Button from '$lib/components/button.svelte';
    import Chips from '$lib/components/chips.svelte';
    import Content from '$lib/components/content.svelte';
    import DataSearch from '$lib/components/data-search.svelte';
    import { confirm } from '$lib/components/dialog.svelte';
    import notification from '$lib/components/notification/notification';
    import Row from '$lib/components/row.svelte';
    import Textfield from '$lib/components/textfield.svelte';
    import Titlebar from '$lib/components/titlebar.svelte';
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';

    let disabled = false;
    let permissionID = 0;
    let data = { id: 0, name: '', description: '' };

    let apiItems = new Array<Chip>(),
        apiProcessing = false,
        newApi: any = null,
        pageItems = new Array<Chip>(),
        pageProcessing = false,
        newPage: any = null;

    function save() {
        if (data.description.trim() == '') {
            notification.error('invalid description');
            return;
        }
        disabled = true;

        if (data.id == 0) {
            api.post('/api/admin/permissions/add', data)
                .then((resp) => {
                    if (resp.status == 0) {
                        goto('/admin/permissions/details?id=' + resp.data);
                        if (resp.data != null) refresh('');
                    } else notification.error(resp.message);
                })
                .catch((e) => notification.error(e))
                .finally(() => (disabled = false));
        } else {
            api.post('/api/admin/permissions/update', data)
                .then((resp) => {
                    if (resp.status == 0)
                        notification.show('Success', 'operation success');
                    else notification.error(resp.message);
                })
                .catch((e) => notification.error(e))
                .finally(() => (disabled = false));
        }
    }

    function remove(ev: CustomEvent, refType: string) {
        console.log(ev.detail);
        confirm(
            'Remove',
            'Delete [' + ev.detail.label + '] from this permission?'
        ).then((answer) => {
            if (answer) {
                let final = () => {};
                switch (refType) {
                    case 'api':
                        apiProcessing = true;
                        final = () => (apiProcessing = false);
                        break;
                    case 'page':
                        pageProcessing = true;
                        final = () => (pageProcessing = false);
                        break;
                }
                api.post('/api/admin/permissions/remove_item', {
                    permission_id: permissionID,
                    ref_id: ev.detail.data.id,
                    ref_type: refType,
                })
                    .then((resp) => {
                        if (resp.status == 0) refresh(refType);
                        else notification.error(resp.message);
                    })
                    .catch((e) => notification.error(e))
                    .finally(() => final());
            }
        });
    }

    //add refType=[api,page]
    function add(refType: string) {
        let refID = 0;
        let final: any = null;
        let success: any = null;
        switch (refType) {
            case 'api':
                if (newApi == null) return;
                refID = newApi.id;
                apiProcessing = true;
                success = () => {
                    newApi = null;
                    refresh('api');
                };
                final = () => (apiProcessing = false);
                break;

            case 'page':
                if (newPage == null) return;
                refID = newPage.id;
                pageProcessing = true;
                success = () => {
                    newPage = null;
                    refresh('page');
                };
                final = () => (pageProcessing = false);
                break;
        }
        api.postData('/api/admin/permissions/add_item', {
            permission_id: permissionID,
            ref_id: refID,
            ref_type: refType,
        })
            .then((data) => {
                notification.show('Success', 'operation success');
                success();
            })
            .catch((e) => notification.error(e))
            .finally(() => {
                if (final != null) final();
            });
    }

    function parseApi(apis: any[] | null) {
        if (apis == null) return;
        apis.forEach((api) => {
            const chip = <Chip>{
                label: api.method + ' ' + api.path,
                icon: api.secure ? 'lock' : '',
                data: api,
            };
            apiItems = [...apiItems, chip];
        });
    }

    function parsePage(pages: any[] | null) {
        if (pages == null) return;
        pages.forEach((page) => {
            const chip = <Chip>{
                label: page.name + ' (' + page.path + ')',
                icon: '',
                data: page,
            };
            pageItems = [...pageItems, chip];
        });
    }

    function refresh(types: string) {
        data = { id: 0, name: '', description: '' };

        api.get(
            '/api/admin/permissions/details?id=' +
                permissionID +
                '&data=' +
                types
        )
            .then((resp) => {
                if (resp.status == 0) {
                    if (types.indexOf('api') >= 0) {
                        apiItems = [];
                        parseApi(resp.data.apis);
                    }
                    if (types.indexOf('page') >= 0) {
                        pageItems = [];
                        parsePage(resp.data.pages);
                    }
                    delete resp.data['apis'], resp.data['pages'];
                    data = resp.data;
                } else notification.error(resp.message);
            })
            .catch((e) => notification.error(e));
    }

    onMount(() => {
        const id = get(page).url.searchParams.get('id');
        if (id != null) {
            permissionID = +id;
            refresh('api,page');
        }
    });
</script>

{#key permissionID}
    <Titlebar>
        {#if permissionID == null}
            Add Permissions
        {:else}Details Permissions{/if}
    </Titlebar>
    <Content>
        <div class="details">
            <Row>
                <Textfield
                    bind:value={data.name}
                    label="Name"
                    placeholder="(Optional)"
                    {disabled}
                />
                <Textfield
                    bind:value={data.description}
                    label="Description"
                    {disabled}
                />
            </Row>
            <Row class="mt-2">
                <div class="flex-1" />
                <Button on:click={save} {disabled}>Save</Button>
            </Row>
        </div>

        <Row>
            <div class="allow-col">
                <h2>Allowed Pages</h2>
                <Chips
                    data={pageItems}
                    on:click={(ev) => remove(ev, 'page')}
                    disabled={pageProcessing}
                    button="clear"
                    vertical
                />
                <Row>
                    <DataSearch
                        placeholder="Add Page"
                        bind:value={newPage}
                        method="POST"
                        url="/api/admin/permissions/add_item?search=[keyword]&search_type=page"
                        disabled={pageProcessing}
                    />
                    {#if newPage != null}
                        <Button
                            on:click={() => add('page')}
                            disabled={pageProcessing}
                        >
                            Add
                        </Button>
                    {/if}
                </Row>
            </div>
            <div class="allow-col">
                <h2>Allowed APIs</h2>
                <Chips
                    data={apiItems}
                    on:click={(ev) => remove(ev, 'api')}
                    disabled={apiProcessing}
                    button="clear"
                    vertical
                />
                <Row>
                    <DataSearch
                        placeholder="Add API"
                        bind:value={newApi}
                        method="POST"
                        url="/api/admin/permissions/add_item?search=[keyword]&search_type=api"
                        disabled={apiProcessing}
                    />
                    {#if newApi != null}
                        <Button
                            on:click={() => add('api')}
                            disabled={apiProcessing}>Add</Button
                        >
                    {/if}
                </Row>
            </div>
        </Row>
    </Content>
{/key}

<style lang="scss">
    .details {
        @apply flex flex-col;
        max-width: 25rem;
    }
    .allow-col {
        @apply w-1/2 flex flex-col p-2 rounded space-y-1 border;
        max-width: 25rem;
        & > :global(div) {
            margin-top: 0.5rem;
        }
    }
    h2 {
        @apply text-lg font-medium;
    }
</style>
