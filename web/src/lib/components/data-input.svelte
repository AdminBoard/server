<script lang="ts">
	import api from '$lib/api';
	import { createEventDispatcher, onMount } from 'svelte';

	export let name = '';
	export let value = '';
	export let width = '16';
	export let disabled = false;
	export let allowCancel = false;

	export let saveUrl = '';
	export let saveCallback: Function | null = null;
	export { clazz as class };

	let dispath = createEventDispatcher();
	let clazz = '';

	let oldValue = '';

	function save() {
		if (saveUrl == '') {
			dispath('save', value);
			value = '';
			oldValue = '';
		} else {
			let params: Record<string, any> = {};
			params[name] = value;

			api.post(saveUrl, params).then((resp) => {
				if (resp.status == 0) {
					saveCallback!(null);
					oldValue = value;
				} else saveCallback!(resp.message);
			});
		}
	}

	onMount(() => {
		oldValue = value;
	});
</script>

<div class="content flex relative items-center w-fit {clazz}" {disabled}>
	<input
		class="pr-2 border"
		style="width:{width}rem"
		type="text"
		class:rounded-none={value != oldValue}
		class:rounded-l={value != oldValue}
		bind:value
		{disabled}
	/>
	<div
		class="flex button-layout rounded-r z-10"
		class:invisible={value == oldValue}
	>
		{#if allowCancel}
			<button
				data-tooltip="cancel"
				class="flex secondary material-icons rounded-none p-2 transition-colors"
				on:click={() => {
					value = '';
					oldValue = '';
					dispath('cancel');
				}}>cancel</button
			>
		{/if}
		<button
			data-tooltip="ok"
			class="secondary material-icons rounded-none rounded-r p-2 transition-colors"
			on:click={save}
			{disabled}>done</button
		>
	</div>
</div>

<style lang="scss">
	@import '../../styles/colors';

	*:disabled {
		filter: opacity(0.5);
	}
	.button-layout {
		position: absolute;
		right: 0;
		transform: translateX(100%);
		box-shadow: 0 0 4px $color-secondary;
	}
</style>
