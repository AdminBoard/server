<script lang="ts">
    import type { Chip } from '$lib/class/chip';
    import { createEventDispatcher } from 'svelte';

    export let data: Chip[] = [];

    const dispatch = createEventDispatcher();
</script>

<div class="chips">
    {#each data as chip}
        <button on:click={() => dispatch('click', chip)}>{chip.label}</button>
    {/each}
</div>

<style lang="scss">
    @import '../../styles/colors';

    .chips {
        @apply space-x-1;
        & button {
            @apply rounded-full px-3 py-1;
        }
    }
</style>
