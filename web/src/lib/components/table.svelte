<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let columns: any[] = [];
	export let rows: any[] = [];

	const dispatch = createEventDispatcher();

	export let selectable = false;

	function formatNumber(number: number) {
		return number.toString().replace(/\B(?=(\d{3})+(?!\d))/g, '.');
	}

	function select(row: any) {
		if (selectable == false) return;
		dispatch('select', row);
	}
</script>

<table
	class="text-sm w-full max-w-full overflow-y-scroll"
	class:selectable
	cellpadding="0"
	cellspacing="0"
>
	<tr>
		{#each columns as col}
			<th>{col.label}</th>
		{/each}
	</tr>
	{#each rows as row}
		<tr on:click={() => select(row)}>
			{#each columns as col}
				{#if row[col.name] != null}
					{#if col.format == 'number'}
						<td class="text-right">{formatNumber(row[col.name])}</td>
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

<div>&nbsp;</div>

<style lang="scss">
	@import '../../styles/colors';

	table.selectable {
		tr:hover {
			@apply cursor-pointer;
			background-color: $color-hover;
		}
	}

	tr {
		border-top: 1px solid $color-border;

		&:last-of-type {
			border-bottom: 1px solid $color-border;
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
