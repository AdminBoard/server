<script lang="ts">
    import { page } from '$app/stores';
    import api from '$lib/api';
    import Button from '$lib/components/button.svelte';
    import Content from '$lib/components/content.svelte';
    import notification from '$lib/components/notification/notification';
    import Textfield from '$lib/components/textfield.svelte';
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';

    let parameter = '',
        property = 'data',
        query = '';

    let apiID = '';

    onMount(() => {
        apiID = get(page).params.slug;
        api.get('/admin/apis/details?with_query=1&id=' + apiID).then((resp) => {
            if (resp.status == 0) {
            }
        });
    });

    function click() {
        api.post('/admin/apis/update?id=' + apiID)
            .then((resp) => {
                if (resp.status == 0) {
                }
            })
            .catch((e) => notification.show('Error', e));
    }
</script>

<Content>
    <div class="add">
        <div>Add Query:</div>
        <div class="flex flex-row space-x-2">
            <Textfield label="Parameter" bind:value={parameter} />
            <Textfield label="Property" bind:value={property} />
        </div>
        <Textfield type="textarea" label="Query" bind:value={query} />
        <div class="flex">
            <div class="flex-1" />
            <Button on:click={click}>Save</Button>
        </div>
    </div>
</Content>

<style lang="scss">
    @import '../../../../../styles/colors';

    .add {
        @apply p-4 rounded shadow;
        border: 1px solid $color-border;
    }
</style>
