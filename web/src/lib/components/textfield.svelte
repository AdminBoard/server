<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import { text } from 'svelte/internal';

	export let label: string = '';
	export let value: string = '';
	export let placeholder: string = '';
	export let disabled: boolean = false;
	export let type = 'text';
	export let resizeable = '';

	export { clazz as class };
	let clazz = '';

	let format: any;
	let textarea: HTMLTextAreaElement;

	const dispatch = createEventDispatcher();

	function input(ev: any) {
		if (type == 'currency') {
			if (format != null) {
				clearTimeout(format);
				format = null;
			}
			format = setTimeout(() => {
				const strip = value.replaceAll('.', '').replaceAll(' ', '');
				const int = parseInt(strip);

				if (!isNaN(int)) {
					value = int.toLocaleString('id-ID');
				}
			}, 1000);
		}
		dispatch('input', { value: value, data: ev.data });
	}

	function keydown(ev: KeyboardEvent) {
		if (ev.code == 'Tab') {
			ev.preventDefault();
			const start = textarea.selectionStart,
				end = textarea.selectionEnd;
			textarea.value =
				textarea.value.substring(0, start) +
				'\t' +
				textarea.value.substring(end);
			textarea.selectionStart = textarea.selectionEnd = start + 1;
		}
	}
</script>

<div class:disabled class:nolabel={label == ''} class="textfield {clazz}">
	{#if type == 'password'}
		<input
			name="input"
			type="password"
			{disabled}
			bind:value
			on:input={(ev) => dispatch('input', ev)}
			on:keyup={(ev) => dispatch('keyup', ev)}
			on:focus={(ev) => dispatch('focus', ev)}
			on:blur={(ev) => dispatch('blur', ev)}
		/>
	{:else if type == 'textarea'}
		<textarea
			bind:this={textarea}
			name="input"
			bind:value
			{disabled}
			on:keydown={keydown}
			style="resize: {resizeable}"
		/>
	{:else}
		<input
			name="input"
			type="text"
			{disabled}
			bind:value
			on:input={input}
			on:keyup={(ev) => dispatch('keyup', ev)}
			on:focus={(ev) => dispatch('focus', ev)}
			on:blur={(ev) => dispatch('blur', ev)}
		/>
	{/if}
	{#if label != ''}
		<label for="input">{label}</label>
	{/if}
</div>
