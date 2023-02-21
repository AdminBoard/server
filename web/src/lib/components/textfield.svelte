<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let label: string = '';
	export let value: string = '';
	export let placeholder: string = '';
	export let disabled: boolean = false;
	export let type = 'text';
	export let resizable = '';
	export let width: string = '';

	export { clazz as class };
	let clazz = '';

	let format: any;

	let inputComp: HTMLInputElement;
	let textareaComp: HTMLTextAreaElement;

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
			const start = textareaComp.selectionStart,
				end = textareaComp.selectionEnd;
			textareaComp.value =
				textareaComp.value.substring(0, start) +
				'\t' +
				textareaComp.value.substring(end);
			textareaComp.selectionStart = textareaComp.selectionEnd = start + 1;
		}
	}

	export function focus() {
		inputComp.focus();
	}
</script>

<div
	class:disabled
	class:nolabel={label == ''}
	class="textfield {clazz}"
	style="width: {width == '' ? 'auto' : width + 'rem'};"
>
	{#if type == 'password'}
		<input
			name="input"
			type="password"
			bind:value
			on:input={(ev) => dispatch('input', ev)}
			on:keyup={(ev) => dispatch('keyup', ev)}
			on:focus={(ev) => dispatch('focus', ev)}
			on:blur={(ev) => dispatch('blur', ev)}
			{placeholder}
			{disabled}
		/>
	{:else if type == 'textarea'}
		<textarea
			bind:this={textareaComp}
			name="input"
			bind:value
			on:keydown={keydown}
			style="resize: {resizable}"
			{placeholder}
			{disabled}
		/>
	{:else}
		<input
			bind:this={inputComp}
			bind:value
			name="input"
			type="text"
			on:input={input}
			on:keyup={(ev) => dispatch('keyup', ev)}
			on:focus={(ev) => dispatch('focus', ev)}
			on:blur={(ev) => dispatch('blur', ev)}
			{placeholder}
			{disabled}
		/>
	{/if}
	{#if label != ''}
		<label for="input">{label}</label>
	{/if}
</div>
