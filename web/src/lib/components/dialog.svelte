<script lang="ts" context="module">
    import { writable, type Writable } from 'svelte/store';
    import Button from './button.svelte';

    export function confirm(title: string, message: string): Promise<boolean> {
        return new Promise<boolean>((resolve) => {
            mode.set('confirm');
            _title = title;
            _message = message;
            confirmAnswer.set(null);
            let unsubscribe = confirmAnswer.subscribe((answer) => {
                if (answer != null) {
                    resolve(answer);
                    unsubscribe();
                }
            });
        });
    }

    let _title = '',
        _message = '';

    let mode: Writable<string> = writable('');
    let confirmAnswer: Writable<boolean | null> = writable(null);
</script>

<script lang="ts">
    import Content from './content.svelte';
    import Row from './row.svelte';
    import Titlebar from './titlebar.svelte';

    export let visible = false;

    function dismiss() {
        visible = false;
        mode.set('');
    }
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
    class="modal"
    class:invisible={!visible && $mode == ''}
    on:click|self={dismiss}
>
    <div class="content">
        {#if $$slots.default}
            <slot />
        {:else if $mode == 'confirm'}
            <Titlebar>{_title}</Titlebar>
            <Content>
                {_message}
                <Row class="mx-2 mt-4">
                    <Button
                        class="flex-auto"
                        on:click={() => {
                            confirmAnswer.set(false);
                            dismiss();
                        }}
                    >
                        No
                    </Button>
                    <Button
                        class="flex-auto"
                        on:click={() => {
                            confirmAnswer.set(true);
                            dismiss();
                        }}
                    >
                        Yes
                    </Button>
                </Row>
            </Content>
        {/if}
    </div>
</div>

<style lang="scss">
    .modal {
        @apply absolute top-0 left-0 w-full h-full z-10;
        background-color: transparentize(#000000, 0.5);
        & .content {
            @apply absolute bg-gray-100 space-y-1 w-fit h-fit drop-shadow-xl rounded overflow-clip;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
            min-width: 400px;
            overflow: visible;
        }
    }
</style>
