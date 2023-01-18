<script lang="ts">
    import api from '$lib/api';
    import Button from '$lib/components/button.svelte';
    import Content from '$lib/components/content.svelte';
    import notification from '$lib/components/notification/notification';
    import Textfield from '$lib/components/textfield.svelte';
    import Titlebar from '$lib/components/titlebar.svelte';

    let method = 'GET';
    let path = '/';
    let isSecure = '0';
    let description = '';

    let disabled = false;

    function click() {
        api.postData('admin/apis/add', {
            method: method,
            url: path,
            is_secure: isSecure,
            description: description,
        })
            .then((_) => {
                notification.show('Success', 'operation completed');
                path = '/';
                description = '';
            })
            .catch((e) => notification.show('Error', e));
    }
</script>

<Titlebar>Add API</Titlebar>
<Content>
    <Textfield label="Method" width="8" bind:value={method} {disabled} />
    <Textfield label="Path" width="24" bind:value={path} {disabled} />
    <Textfield label="Is Secure" width="8" bind:value={isSecure} {disabled} />
    <Textfield
        label="Description"
        type="textarea"
        width="24"
        bind:value={description}
        {disabled}
    />

    <Button on:click={click} {disabled}>Save</Button>
</Content>
