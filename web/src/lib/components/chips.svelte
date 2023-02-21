<script lang="ts">
    import type { Chip } from '$lib/class/chip';
    import { createEventDispatcher } from 'svelte';
    import IconButton from './icon-button.svelte';

    export let data: Chip[] = [];
    export let vertical = false;
    export let disabled = false;
    export let button = '';
    // export let removable = false;

    const dispatch = createEventDispatcher();
</script>

<!-- <div class="chips" class:vertical>
    {#each data as chip}
        <div class="chip" class:removable>
            {#if chip.icon != null && chip.icon != ''}
                <i class="icon material-icons">{chip.icon}</i>
            {/if}
            {chip.label}
            {#if removable}
                <button
                    class="material-icons"
                    on:click={() => dispatch('remove', chip)}
                    {disabled}>close</button
                >
            {/if}
        </div>
    {/each}
</div> -->

<div>
    <div class="chips" class:vertical class:disabled>
        {#each data as chip}
            <span class="chip" class:hasButton={button != ''}>
                {#if chip.icon != null && chip.icon != ''}
                    <i class="icon material-icons">{chip.icon}</i>
                {/if}
                {chip.label}
                {#if button != ''}
                    <IconButton
                        on:click={() => dispatch('click', chip)}
                        bind:disabled
                    >
                        {button}
                    </IconButton>
                {/if}
            </span>
        {/each}
    </div>
</div>

<style lang="scss">
    @import '../../styles/colors';

    .chips {
        &.vertical {
            & > .chip {
                @apply block w-fit;
            }
        }
        & > .chip {
            @apply inline-block m-0.5 px-2 py-0 rounded-full leading-10;
            background-color: $color-border;

            &.hasButton {
                @apply pr-1;
            }

            & > :global(button) {
                @apply text-sm py-0.5 px-1 rounded-full;
                background-color: transparent;
                color: $color-secondary;
                &:disabled {
                    color: darken($color-border, 16);
                }
                &:not(:disabled) {
                    &:hover {
                        background-color: transparentize($color-secondary, 0.5);
                        color: $color-text;
                    }
                    &:active {
                        background-color: $color-secondary;
                        color: $color-text;
                    }
                }
            }

            // & i.icon {
            //         @apply text-sm;
            //     }
        }
    }

    // .chips {
    //     @apply flex space-x-1;

    //     &.vertical {
    //         @apply flex-col space-y-1;
    //         & > div {
    //             @apply w-fit;
    //         }
    //     }
    //     & > .chip {
    //         @apply flex rounded-full px-3 py-1;
    //         background-color: $color-border;

    //         & i.icon {
    //             @apply text-sm;
    //         }

    //         &.removable {
    //             @apply pr-1;
    //         }
    //         & > button {
    //             @apply text-sm m-0.5 py-0.5 px-1 rounded-full shadow;
    //             background-color: transparent;
    //             color: $color-secondary;
    //             &:disabled {
    //                 @apply shadow-none;
    //                 color: darken($color-border, 16);
    //             }
    //             &:not(:disabled) {
    //                 &:hover {
    //                     background-color: transparentize($color-secondary, 0.5);
    //                     color: $color-text;
    //                 }
    //                 &:active {
    //                     background-color: $color-secondary;
    //                     color: $color-text;
    //                 }
    //             }
    //         }
    //     }
    // }
</style>
