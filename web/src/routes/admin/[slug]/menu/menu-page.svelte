<script lang="ts">
	import Content from '$lib/components/content.svelte';
	import Titlebar from '$lib/components/titlebar.svelte';
	import MenuDetails from './menu-details.svelte';
	import MenuList from './menu-list.svelte';

	let list: MenuList;
</script>

<Titlebar caption="Menu" />
<Content>
	<div class="flex space-x-2">
		<div>
			<MenuList bind:this={list} />
		</div>
		<div>
			<MenuDetails on:change={() => list?.refresh()} />
		</div>
	</div>
</Content>
