<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let icon = '';
	export let submenu = false;
	export let name = '';
	export let selected = false;
	export let url = '';

	let clazz = '';

	const dispatch = createEventDispatcher();

	export { clazz as class };
</script>

<a
	href={url}
	class:selected
	class:link={url != ''}
	class:sub={submenu}
	class="menu flex space-x-2 transition-colors items-center py-2 px-4 rounded-lg {clazz}"
	on:click={() => {
		dispatch('click');
	}}
>
	{#if icon != ''}
		<span class="material-icons">{icon}</span>
	{/if}
	<span>{name}</span>
	<slot />
</a>

<style lang="scss">
	.menu {
		margin-left: 0.2rem;
		margin-right: 0.2rem;
		cursor: default;
		&.sub {
			padding-left: 2.5rem;
		}
		&.selected {
			background-color: transparentize(#000, 0.8);
		}
		&.link:hover {
			cursor: pointer;
			background-color: transparentize(#000, 0.9);
		}
	}
</style>
