<script lang="ts">
	import components from '$lib/components';
	import { createEventDispatcher } from 'svelte';
	import type { LayoutNode } from './layout-node';

	export let data: Record<string, any>;

	export let node: LayoutNode;

	export let value: any | null = null;

	const dispatch = createEventDispatcher();

	function click(ev: any) {
		dispatch('click', ev.detail);
	}

	function renderText(text: string): string {
		const regex = /\{[._a-z]+\}/gi;

		const matches = text.match(regex);
		if (matches != null) {
			matches.forEach((match) => {
				const varName = match.substring(1, match.length - 1);
				const split = varName.split('.');

				let dataPtr = data;

				for (const i in split) {
					const key = split[i];

					const type = typeof dataPtr[key];

					if (type == 'string' || type == 'number') {
						text = text.replaceAll('{' + key + '}', dataPtr[key]);
					} else {
						if (dataPtr[key] != null) dataPtr = dataPtr[key];
						else text = text.replaceAll('{' + key + '}', '');
					}
				}
			});
		}
		return text;
	}
</script>

{#if node.tag != null}
	<svelte:component
		this={components.get(node.tag)}
		bind:value={data[node.name]}
		{...node.attrs}
		on:click={click}
	>
		{#if node.children != null}
			{#each node.children as subchild}
				<svelte:self {data} on:click={click} node={subchild} />
			{/each}
		{:else}
			{node.text}
		{/if}
	</svelte:component>
{:else}
	<!-- {@html node.text} -->
	{@html renderText(node.text)}
{/if}
