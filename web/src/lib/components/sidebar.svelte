<script lang="ts">
	import session from '$lib/session';
	import Menu from './menu.svelte';

	let menu: any[] = [];

	let showSidebar = false;

	session.menu.subscribe((m) => (menu = m));
</script>

<input
	type="checkbox"
	id="menu-open"
	bind:checked={showSidebar}
	class="hidden"
/>

<header class="bg-primary flex justify-between md:hidden">
	<a
		href="/dashboard"
		class="block p-4 text-white font-bold whitespace-nowrap truncate"
	>
		Adminboard
	</a>

	<label
		for="menu-open"
		id="mobile-menu-button"
		class="m-2 p-2 focus:outline-none hover:text-white hover:bg-gray-700 rounded-md"
	>
		<svg
			id="menu-open-icon"
			class="h-6 w-6 transition duration-200 ease-in-out"
			xmlns="http://www.w3.org/2000/svg"
			fill="none"
			viewBox="0 0 24 24"
			stroke="currentColor"
		>
			<path
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
				d="M4 6h16M4 12h16M4 18h16"
			/>
		</svg>
		<svg
			id="menu-close-icon"
			class="h-6 w-6 transition duration-200 ease-in-out"
			xmlns="http://www.w3.org/2000/svg"
			fill="none"
			viewBox="0 0 24 24"
			stroke="currentColor"
		>
			<path
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
				d="M6 18L18 6M6 6l12 12"
			/>
		</svg>
	</label>
</header>

<aside
	id="sidebar"
	class="bg-primary z-10 md:w-64 w-64 space-y-6 pt-6 px-0 fixed inset-y-0 left-0 transform md:fixed md:translate-x-0 transition duration-200 ease-in-out  md:flex md:flex-col md:justify-between overflow-y-auto"
>
	<div
		class="flex flex-col space-y-6"
		data-dev-hint="optional div for having an extra footer navigation"
	>
		<a
			href="/dashboard"
			class="text-white flex items-center space-x-2 px-4"
			title="Adminboard"
		>
			<img src="/ic_logo.png" width="32" height="32" alt="logo" />
			<span class="text-2xl font-extrabold whitespace-nowrap truncate"
				>Adminboard</span
			>
		</a>

		<nav data-dev-hint="main navigation">
			{#each menu as m}
				<Menu
					icon={m.icon}
					url={m.path}
					on:click={() => (showSidebar = false)}>{m.name}</Menu
				>
				{#if m.children != null}
					{#each m.children as sub}
						<Menu
							submenu={true}
							icon={sub.icon}
							url={sub.path}
							on:click={() => (showSidebar = false)}
							>{sub.name}</Menu
						>
					{/each}
				{/if}
			{/each}
		</nav>
	</div>
</aside>

<style lang="scss">
	#sidebar {
		--tw-translate-x: -100%;
		box-shadow: 0 0 8px transparentize(#000000, 0.5);
	}
	@media (min-width: 768px) {
		#sidebar {
			--tw-translate-x: 0;
			box-shadow: none;
		}
	}
	#menu-close-icon {
		display: none;
	}

	#menu-open:checked ~ #sidebar {
		--tw-translate-x: 0;
	}
	#menu-open:checked ~ * #mobile-menu-button {
		background-color: rgba(31, 41, 55, var(--tw-bg-opacity));
	}
	#menu-open:checked ~ * #menu-open-icon {
		display: none;
	}
	#menu-open:checked ~ * #menu-close-icon {
		display: block;
	}
</style>
