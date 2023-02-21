<script lang="ts">
    import { page } from '$app/stores';
    import api from '$lib/api';
    import type { Chip } from '$lib/class/chip';
    import Chips from '$lib/components/chips.svelte';
    import Content from '$lib/components/content.svelte';
    import { confirm } from '$lib/components/dialog.svelte';
    import notification from '$lib/components/notification/notification';
    import Titlebar from '$lib/components/titlebar.svelte';
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';

    let id: string | null = '',
        name = '',
        status = '';

    let disabled = false;

    let value: any = null;

    let allowed: Chip[] = [],
        denied: Chip[] = [];

    function deny(ev: any) {
        confirm(
            'Remove Permissions',
            'Are you sure to remove this permission?'
        ).then((yes) => {
            if (yes) {
                console.log(ev.detail);

                disabled = true;
                api.postData('/api/admin/groups/remove_permission', {
                    permission_id: ev.detail.data.id,
                    group_id: id,
                })
                    .then(() => refresh())
                    .finally(() => (disabled = false));
            }
        });
    }

    function allow(ev: any) {
        confirm('Add Permissions', 'Are you sure to add this permission?').then(
            (yes) => {
                if (yes) {
                    disabled = true;
                    api.postData('/api/admin/groups/add_permission', {
                        permission_id: ev.detail.data.id,
                        group_id: id,
                    })
                        .then(() => refresh())
                        .catch((e) => notification.error(e))
                        .finally(() => (disabled = false));
                }
            }
        );
    }

    function refresh() {
        api.getData('/api/admin/groups/details?id=' + id)
            .then((data) => {
                id = data.id;
                name = data.name;
                status = data.status;
                if (data.permissions != null) {
                    allowed = [];
                    denied = [];
                    data.permissions.forEach((perm: any) => {
                        const label =
                            perm.name == null || perm.name == ''
                                ? perm.description
                                : perm.name + ' (' + perm.description + ')';

                        if (perm.group_id == null) {
                            denied = [
                                ...denied,
                                {
                                    label: label,
                                    data: perm,
                                },
                            ];
                        } else {
                            allowed = [
                                ...allowed,
                                {
                                    label: label,
                                    data: perm,
                                },
                            ];
                        }
                    });
                }
            })
            .catch((e) => notification.error(e));
    }
    onMount(() => {
        const paramID = get(page).url.searchParams.get('id');
        if (paramID == null) {
            notification.error('invalid id');
            return;
        }

        id = paramID;
        refresh();
    });
</script>

{#if id != null}
    <Titlebar>Group: {name}</Titlebar>
    <Content>
        <div>
            Status: {status}
        </div>
        <div>Permissions Allowed</div>
        <Chips
            bind:data={allowed}
            bind:disabled
            button="clear"
            on:click={deny}
        />
        <div>Permissions Available</div>
        <Chips bind:data={denied} bind:disabled button="add" on:click={allow} />
    </Content>
{/if}
