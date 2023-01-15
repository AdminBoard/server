<script lang="ts">
	import { getOverride } from '$lib/http';
	import Textfield from './textfield.svelte';

	export let label = '';
	export let value: any | null = null;

	let textValue = '';
	let focus = false;
	let process: NodeJS.Timeout | null = null;

	let abortController: AbortController | null;

	let items: any[] = [];
	let focusIndex = 0;

	function input() {
		items = [];
		focusIndex = 0;
		focus = false;
		value = null;
		if (process != null) clearTimeout(process);
		if (abortController != null) abortController.abort();

		abortController = new AbortController();
		process = setTimeout(() => {
			getOverride(
				'/api/deposit/add/search?name=' + textValue,
				abortController?.signal
			).then((resp) => {
				if (resp.status == 0) {
					items = resp.data;
					focus = true;
				}
				// else notification.show('Error', resp.message);
			});
		}, 200);
	}

	function selectItem(item: any) {
		value = item;
		textValue = value.name;
		focus = false;
		items = [];
	}

	function keyup(ev: any) {
		if (items.length == 0) return;
		switch (ev.detail.code) {
			case 'ArrowDown':
				focusIndex++;
				if (focusIndex == items.length) focusIndex = 0;
				break;
			case 'ArrowUp':
				focusIndex--;
				if (focusIndex == -1) focusIndex = items.length - 1;
				break;
			case 'Enter':
				selectItem(items[focusIndex]);
				break;
		}
	}
</script>

<div class="layout">
	<Textfield
		{label}
		on:input={input}
		bind:value={textValue}
		on:focus={() => (focus = true)}
		on:keyup={keyup}
	/>
	{#if items.length > 0}
		<ul class="list" class:invisible={!focus}>
			{#each items as item, i}
				<li
					on:click|self={() => selectItem(item)}
					class:focus={focusIndex == i}
				>
					{item.name} (Rp{item.balance.toLocaleString('id-ID')})
				</li>
			{/each}
		</ul>
	{/if}
</div>

<style lang="scss">
	@import '../../styles/colors';

	ul.list {
		@apply absolute z-10 shadow overflow-y-auto max-h-60 border-t-0 border-x-2 border-b-2 left-3 right-3 cursor-pointer;
		transform: translateY(-0.1rem);
		background-color: $color-text;
		& > * {
			@apply px-2 py-1;
			&.focus {
				background-color: $color-hover;
			}
		}
	}
</style>
