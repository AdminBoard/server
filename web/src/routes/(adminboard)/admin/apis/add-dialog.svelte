<script lang="ts">
    import api from '$lib/api';
    import Button from '$lib/components/button.svelte';
    import Content from '$lib/components/content.svelte';
    import Dialog from '$lib/components/dialog.svelte';
    import Textfield from '$lib/components/textfield.svelte';
    import Titlebar from '$lib/components/titlebar.svelte';
    import { createEventDispatcher } from 'svelte';

    let visible = false;
    let disabled = false;

    const dispatch = createEventDispatcher();

    let method = 'GET',
        path = '/',
        isSecure = '0',
        description = '';

    export function dismiss() {
        visible = false;
    }

    export function show() {
        disabled = false;
        method = 'GET';
        path = '/';
        isSecure = '0';
        description = '';
        api.get('/api/page?path=/admin/apis/add').then((resp) => {
            if (resp.status == 0) visible = true;
            else dispatch('error', resp.message);
        });
    }

    function save() {
        disabled = true;
        api.postData('admin/apis/add', {
            method: method,
            path: path,
            secure: isSecure,
            description: description,
        })
            .then(() => dispatch('success'))
            .catch((e) => dispatch('error', e));
    }
</script>

<Dialog bind:visible>
    <Titlebar>Add API</Titlebar>
    <Content>
        <Textfield label="Method" width="8" bind:value={method} bind:disabled />
        <Textfield label="Path" width="24" bind:value={path} bind:disabled />
        Secure
        <Textfield
            label="Is Secure"
            width="8"
            bind:value={isSecure}
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
