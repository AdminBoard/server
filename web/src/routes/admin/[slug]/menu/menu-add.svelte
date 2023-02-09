<script lang="ts">
    import { goto } from '$app/navigation';
    import api from '$lib/api';
    import notification from '$lib/components/notification/notification';
    import Textfield from '$lib/components/textfield.svelte';
    import { createEventDispatcher } from 'svelte';

    export let parentID = 0;
    export let tooltip = '';

    const dispatch = createEventDispatcher();

    let value = '';
    let textfield: Textfield;
    let typing = false;
    let disabled = false;

    function click() {
        typing = true;
        setInterval(() => textfield.focus(), 100);
    }

    function done() {
        disabled = true;
        api.post('admin/menu/add', { parent_id: parentID, name: value })
            .then((resp) => {
                if (resp.status == 0) {
                    const data = resp.data;
                    if (data != null) {
                        goto('/admin/menu#id=' + data.id);
                    }
                    typing = false;
                    value = '';
                    dispatch('success');
                } else notification.show('Error', resp.message);
            })
            .catch((e) => notification.show('Error', e))
            .finally(() => (disabled = false));
    }
</script>

<div class="add-menu">
    <button
        data-tooltip={tooltip}
        class:invisible={typing}
        class="tooltip-right material-icons"
        on:click={click}
        {disabled}
    >
        add
    </button>
    <div class="confirm" class:invisible={!typing}>
        <Textfield bind:this={textfield} bind:value {disabled} />
        <div class="buttons">
            <button
                class="material-icons"
                on:click={() => (typing = false)}
                {disabled}>undo</button
            >
            <button class="material-icons" on:click={done} {disabled}
                >done</button
            >
        </div>
    </div>
</div>

<style lang="scss">
    @import '../../../../styles/colors';

    .add-menu {
        @apply relative w-full items-center;

        & > button {
            @apply absolute cursor-pointer rounded-full flex px-3 py-2 m-1 top-0 bottom-0 text-sm;
            background-color: transparent;
            color: $color-primary;
        }
    }
    .confirm {
        @apply relative;

        & .buttons {
            @apply absolute top-0 right-0 h-full translate-x-full rounded-r;
            background-color: $color-border;
            transform: translateX(calc(100% - 8px));
        }
        & button {
            @apply h-full text-sm px-4 rounded-none;
            background-color: transparent;
            color: $color-primary;
        }
    }
</style>
