<script lang="ts">
    import { goto } from '$app/navigation';
    import api from '$lib/api';
    import Button from '$lib/components/button.svelte';
    import Content from '$lib/components/content.svelte';
    import DataTable from '$lib/components/data-table.svelte';
    import Dialog from '$lib/components/dialog.svelte';
    import IconButton from '$lib/components/icon-button.svelte';
    import notification from '$lib/components/notification/notification';
    import Row from '$lib/components/row.svelte';
    import Textfield from '$lib/components/textfield.svelte';
    import Titlebar from '$lib/components/titlebar.svelte';

    //form fields
    let name = '';
    let visible = false; //add form visibility
    let refreshCounter = 0;

    function dismiss() {
        name = '';
        visible = false;
    }

    function popupAddForm() {
        dismiss();
        api.get('/api/page?path=/admin/groups/add').then((resp) => {
            if (resp.status == 0) visible = true;
            else notification.error(resp.message);
        });
    }

    function save() {
        api.postData('/api/admin/groups/add', { name: name })
            .then(() => {
                notification.show('Success', 'Successfully add group');
                dismiss();
                refreshCounter++;
            })
            .catch((e) => notification.error(e));
    }
</script>

<Titlebar>
    <div>Groups</div>
    <IconButton on:click={popupAddForm}>add</IconButton>
</Titlebar>
{#key refreshCounter}
    <DataTable
        src="/api/admin/groups"
        columns={[
            { name: 'id', label: 'ID', format: 'right' },
            { name: 'name', label: 'Nama' },
            { name: 'status', label: 'Status' },
        ]}
        on:select={(ev) => goto('/admin/groups/details?id=' + ev.detail.id)}
        selectable
    />
{/key}

<!-- Add Form -->
<Dialog bind:visible>
    <Titlebar>Add Group</Titlebar>
    <Content>
        <Textfield bind:value={name} label="Group name" />
        <Row>
            <div class="flex-1" />
            <Button on:click={dismiss}>Cancel</Button>
            <Button on:click={save}>Save</Button>
        </Row>
    </Content>
</Dialog>
