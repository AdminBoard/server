<script lang="ts">
    import { goto } from '$app/navigation';
    import DataTable from '$lib/components/data-table.svelte';
    import IconButton from '$lib/components/icon-button.svelte';
    import notification from '$lib/components/notification/notification';
    import Titlebar from '$lib/components/titlebar.svelte';
    import AddDialog from './add-dialog.svelte';

    let addDialog: AddDialog;
    let table: DataTable;

    function success() {
        notification.show('Success', 'operation completed');
        addDialog.dismiss();
        table.refresh();
    }
</script>

<Titlebar>
    <div>APIs</div>
    <IconButton on:click={() => addDialog.show()}>add</IconButton>
</Titlebar>
<DataTable
    bind:this={table}
    src="/api/admin/apis"
    columns={[
        { name: 'path', label: 'Path' },
        { name: 'method', label: 'Method' },
        { name: 'secure', label: 'Secure' },
        { name: 'description', label: 'Description' },
        { name: 'status', label: 'Status' },
    ]}
    on:error={(ev) => notification.error(ev.detail)}
    on:select={(ev) =>
        goto(
            '/admin/apis/details?path=' +
                ev.detail.path +
                '&method=' +
                ev.detail.method
        )}
    selectable
/>

<AddDialog
    bind:this={addDialog}
    on:success={success}
    on:error={(ev) => notification.error(ev.detail)}
/>
