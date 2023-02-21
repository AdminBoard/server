<script lang="ts">
    import { goto } from '$app/navigation';
    import api from '$lib/api';
    import Button from '$lib/components/button.svelte';
    import Content from '$lib/components/content.svelte';
    import DataTable from '$lib/components/data-table.svelte';
    import Dialog from '$lib/components/dialog.svelte';
    import IconButton from '$lib/components/icon-button.svelte';
    import notification from '$lib/components/notification/notification';
    import Textfield from '$lib/components/textfield.svelte';
    import Titlebar from '$lib/components/titlebar.svelte';

    //form fields
    let method = '',
        path = '',
        is_secure = '',
        description = '';

    let visible = false; //add form visibility
    let disabled = false;

    function dismiss() {
        method = '';
        path = '';
        is_secure = '';
        description = '';
        visible = false;
    }

    function popupAddForm() {
        dismiss();
        api.get('/api/page?path=/admin/apis/add').then((resp) => {
            if (resp.status == 0) {
                visible = true;
            } else notification.error(resp.message);
        });
    }

    function save() {
        disabled = true;
        api.postData('admin/apis/add', {
            method: method,
            path: path,
            is_secure: is_secure,
            description: description,
        })
            .then(() => {
                notification.show('Success', 'operation completed');
                dismiss();
            })
            .catch((e) => notification.error(e));
    }
</script>

<Titlebar>
    <div>APIs</div>
    <IconButton on:click={popupAddForm}>add</IconButton>
</Titlebar>
<DataTable
    src="/api/admin/apis"
    columns={[
        { name: 'path', label: 'Path' },
        { name: 'method', label: 'Method' },
        { name: 'secure', label: 'Secure' },
        { name: 'description', label: 'Description' },
        { name: 'status', label: 'Status' },
    ]}
    on:select={(ev) =>
        goto(
            '/admin/apis/details?path=' +
                ev.detail.path +
                '&method=' +
                ev.detail.method
        )}
    selectable
/>

<!-- Add Form -->
<Dialog bind:visible>
    <Titlebar>Add API</Titlebar>
    <Content>
        <Textfield label="Method" width="8" bind:value={method} bind:disabled />
        <Textfield label="Path" width="24" bind:value={path} bind:disabled />
        <Textfield
            label="Is Secure"
            width="8"
            bind:value={is_secure}
            bind:disabled
        />
        <Textfield
            label="Description"
            type="textarea"
            width="24"
            bind:value={description}
            bind:disabled
        />
        <Button on:click={save} bind:disabled>Save</Button>
    </Content>
</Dialog>
