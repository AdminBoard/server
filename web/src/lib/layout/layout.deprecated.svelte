<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import type { LayoutNode } from './layout-node';
	import Renderer from './renderer.svelte';

	export let layout: LayoutNode | null = null;

	export let data: Record<string, any> = {};

	const dispath = createEventDispatcher();

	function click(ev: any) {
		dispath('click', ev.detail);
	}
</script>

{#if layout != null && layout.children != null}
	{#each layout.children as node}
		{#if node.tag != null}
			<Renderer {data} on:click={click} {node} />
		{:else}
			{@html node.text}
		{/if}
	{/each}
{/if}
