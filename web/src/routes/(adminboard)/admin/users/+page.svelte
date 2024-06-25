<script lang="ts">
    import api from '$lib/api';
    import DataTable from '$lib/components/data-table.svelte';
    import IconButton from '$lib/components/icon-button.svelte';
    import notification from '$lib/components/notification/notification';
    import Titlebar from '$lib/components/titlebar.svelte';

    //form fields
    let name = '',
        username = '';
    let visible = false; //add form visibility
    let refreshCounter = 0;

    function dismiss() {
        name = '';
        username = '';
        visible = false;
    }

    function popupAddForm() {
        dismiss();
        api.get('/api/page?path=/admin/users/add').then((resp) => {
            if (resp.status == 0) visible = true;
            else notification.error(resp.message);
        });
    }
</script>

<Titlebar>
    <div>Users</div>
    <IconButton on:click={popupAddForm}>add</IconButton>
</Titlebar>
{#key refreshCounter}
    <DataTable
        src="/admin/users"
        columns={[
            { name: 'id', label: 'ID' },
            { name: 'name', label: 'Name' },
            { name: 'username', label: 'Username' },
        ]}
    />
{/key}
