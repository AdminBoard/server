<script lang="ts">
    import { goto } from '$app/navigation';
    import api from '$lib/api';
    import DataTable from '$lib/components/data-table.svelte';
    import IconButton from '$lib/components/icon-button.svelte';
    import notification from '$lib/components/notification/notification';
    import Titlebar from '$lib/components/titlebar.svelte';
    import AddDialog from './add-dialog.svelte';

    let visible = false; //add form visibility
    let disabled = false;

    let addDialog: AddDialog;

    function save(ev: any) {
        disabled = true;
        api.postData('/api/admin/pages/add', ev.detail)
            .then(() => {
                notification.show('Success', 'Page successfully added');
                addDialog.dismiss();
            })
            .catch((e) => notification.error(e))
            .finally(() => (disabled = false));
    }
</script>

<Titlebar
    ><span>Pages</span>
    <IconButton on:click={() => addDialog.show()}>add</IconButton>
</Titlebar>

<DataTable
    src="/api/admin/pages"
    columns={[
        { name: 'path', label: 'Path' },
        { name: 'name', label: 'Name' },
        { name: 'description', label: 'Description' },
    ]}
    on:select={(ev) => goto('/admin/pages/details?path=' + ev.detail.path)}
    selectable
/>

<AddDialog bind:this={addDialog} on:save={save} bind:disabled />
