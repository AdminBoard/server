<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let columns: any[] = [];
	export let data: any[] = [];
	export let selectable = false;
	export let multiple = false;

	export let selections: any[] = [];
	let selectionIndexes = new Map<number, boolean>();

	const dispatch = createEventDispatcher();

	function formatNumber(number: number) {
		return number.toString().replace(/\B(?=(\d{3})+(?!\d))/g, '.');
	}

	function refreshSelection() {
		let sels = [];

		for (let key of selectionIndexes.keys()) {
			if (data[key] != null) sels.push(data[key]);
		}
		selections = sels;
		selectionIndexes = selectionIndexes;
	}

	function select(idx: number) {
		if (!selectable && !multiple) return;
		if (selectionIndexes.has(idx)) selectionIndexes.delete(idx);
		else selectionIndexes.set(idx, true);
		if (multiple) {
			refreshSelection();
			dispatch('select', selections);
		} else dispatch('select', data[idx]);
	}
</script>

<table
	class="text-sm w-full max-w-full overflow-y-scroll"
	class:selectable={selectable || multiple}
	cellpadding="0"
	cellspacing="0"
>
	<tr>
		{#each columns as col}
			<th>{col.label}</th>
		{/each}
	</tr>
	{#each data as row, i}
		<tr on:click={() => select(i)} class:selected={selectionIndexes.get(i)}>
			{#each columns as col}
				{#if row[col.name] != null}
					{#if col.format == 'number'}
						<td class="text-right">{formatNumber(row[col.name])}</td
						>
					{:else if col.format == 'right'}
						<td class="text-right">{row[col.name]}</td>
					{:else}
						<td>{row[col.name]}</td>
					{/if}
				{:else}
					<td />
				{/if}
			{/each}
		</tr>
	{/each}
</table>
{#if data.length == 0}
	<div class="w-full text-center pt-2">No Data</div>
{/if}

<style lang="scss">
	@import '../../styles/colors';

	table.selectable {
		tr:has(td):hover {
			@apply cursor-pointer;
			background-color: $color-hover;
		}
	}

	tr {
		border-top: 1px solid $color-border;

		&:last-of-type {
			border-bottom: 1px solid $color-border;
		}

		&.selected {
			background-color: darken($color-border, 10);
		}
	}
	td,
	th {
		padding: 0.3rem;
		border-left: 1px solid $color-border;
		&:last-of-type {
			border-right: 1px solid $color-border;
		}
	}
</style>
