<script lang="ts">
	import { getOverride, postOverride } from '$lib/http';
	import type { SearchResult } from './data-search';
	import Textfield from './textfield.svelte';

	export let label = '';
	export let placeholder = '';
	export let width = '16';
	export let value: any | null = null;
	export let url = '';
	export let method = 'GET';
	export let disabled = false;

	let focus = false;
	let process: NodeJS.Timeout | null = null;
	let textValue = '';

	let abortController: AbortController | null;

	let items: SearchResult[] = [];
	let focusIndex = 0;

	function input() {
		if (url.trim() == '') return;
		items = [];
		focusIndex = 0;
		focus = false;
		value = null;
		if (process != null) clearTimeout(process);
		if (abortController != null) abortController.abort();

		abortController = new AbortController();
		process = setTimeout(() => {
			let reqUrl = url.replaceAll('[keyword]', textValue);
			switch (method) {
				case 'POST':
					postOverride(reqUrl, abortController?.signal).then(
						(resp) => {
							if (resp.status == 0) {
								items = resp.data;
								focus = true;
							}
						}
					);
					break;
				case 'GET':
				default:
					getOverride(reqUrl, abortController?.signal).then(
						(resp) => {
							if (resp.status == 0) {
								items = resp.data;
								focus = true;
							}
						}
					);
			}
		}, 200);
	}

	function selectItem(item: any) {
		value = item;
		textValue = value.label;
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

<div class="data-search" style="width: {width}rem;">
	<Textfield
		{label}
		bind:placeholder
		bind:value={textValue}
		on:input={input}
		on:focus={() => (focus = true)}
		on:blur={() => setTimeout(() => (focus = false), 200)}
		on:keyup={keyup}
		bind:disabled
	/>
	{#if items.length > 0}
		<ul class="list" class:invisible={!focus}>
			{#each items as item, i}
				<li
					on:click={() => selectItem(item)}
					class:focus={focusIndex == i}
					on:keypress
				>
					<div>
						{item.label != null ? item.label : ''}
					</div>
					{#if item.sublabel != null}
						<div class="sub">{item.sublabel}</div>
					{/if}
				</li>
			{/each}
		</ul>
	{/if}
</div>

<style lang="scss">
	@import '../../styles/colors';
	.data-search {
		@apply relative;
	}

	ul.list {
		@apply absolute z-10 shadow overflow-y-auto max-h-60 border-t-0 border-x-2 border-b-2 left-1 right-1 cursor-pointer transition-opacity;
		transform: translateY(-0.1rem);
		background-color: $color-text;

		&.invisible {
			@apply opacity-0 transition-opacity;
		}
		& > li {
			@apply px-2 py-1;
			&.focus {
				background-color: $color-hover;
			}
			& > div.sub {
				@apply text-sm;
				color: $color-gray;
			}
		}
	}
</style>
