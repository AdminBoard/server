<script lang="ts">
	import session from '$lib/session';

	let menu: any[] = [];
	session.menu.subscribe((m) => (menu = m));
</script>

<div class="relative min-h-screen md:flex" data-dev-hint="container">
	<input type="checkbox" id="menu-open" class="hidden" />

	<label
		for="menu-open"
		class="absolute right-2 bottom-2 shadow-lg rounded-full p-2 bg-gray-100 text-gray-600 md:hidden"
		data-dev-hint="floating action button"
	>
		<svg
			class="h-6 w-6"
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
	</label>

	<header
		class="bg-gray-600 text-gray-100 flex justify-between md:hidden"
		data-dev-hint="mobile menu bar"
	>
		<a href="/dashboard" class="block p-4 text-white font-bold whitespace-nowrap truncate">
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
		class="bg-gray-700 text-gray-100 md:w-64 w-3/4 space-y-6 pt-6 px-0 absolute inset-y-0 left-0 transform md:relative md:translate-x-0 transition duration-200 ease-in-out  md:flex md:flex-col md:justify-between overflow-y-auto"
		data-dev-hint="sidebar; px-0 for frameless; px-2 for visually inset the navigation"
	>
		<div
			class="flex flex-col space-y-6"
			data-dev-hint="optional div for having an extra footer navigation"
		>
			<a href="/dashboard" class="text-white flex items-center space-x-2 px-4" title="Adminboard">
				<img src="/ic_logo.png" width="32" height="32" alt="logo" />
				<span class="text-2xl font-extrabold whitespace-nowrap truncate">Adminboard</span>
			</a>

			<nav data-dev-hint="main navigation">
				{#each menu as m}
					<a
						href={m.url}
						class="flex cursor-pointer items-center space-x-1 py-2 px-4 transition duration-200 hover:bg-gray-600 hover:text-white"
					>
						{#if m.icon != ''}
							<span class="material-icons">{m.icon}</span>
						{/if}
						<span>{m.name}</span>
					</a>
					{#each m.children as sub}
						<a
							href={sub.url}
							class="flex ml-4 cursor-pointer items-center space-x-1 py-2 px-4 transition duration-200 hover:bg-gray-600 hover:text-white"
						>
							{#if sub.icon != ''}
								<span class="material-icons">{sub.icon}</span>
							{/if}
							<span>{sub.name}</span>
						</a>
					{/each}
				{/each}
				<!-- <a
					href="/invoice"
					class="flex items-center space-x-2 py-2 px-4 transition duration-200 hover:bg-gray-600 hover:text-white"
				>
					<span class="material-icons">face</span>
					<span>Invoice</span>
				</a> -->
				<!-- <a href="#">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="h-4 w-4"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
					>
						<path d="M12 14l9-5-9-5-9 5 9 5z" />
						<path
							d="M12 14l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14z"
						/>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M12 14l9-5-9-5-9 5 9 5zm0 0l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14zm-4 6v-7.5l4-2.222"
						/>
					</svg>
					<span>About</span>
				</a>
				<a
					href="#"
					class="flex items-center space-x-2 py-2 px-4 transition duration-200 hover:bg-gray-700 hover:text-white"
				>
					<span class="ml-6">Without Icon</span>
				</a> -->
				<!-- <a
					href="#"
					class="flex items-center space-x-2 py-2 px-4 transition duration-200 hover:bg-gray-700 hover:text-white group"
				>
					<span
						class="w-4 h-4 flex-shrink-0 border border-gray-600 rounded group-hover:border-gray-400 transition duration-200"
					/>
					<span>Without Icon And a bit longer than usual</span>
				</a> -->
			</nav>
		</div>

		<!-- <nav data-dev-hint="second-main-navigation or footer navigation">
			<a
				href="#"
				class="block py-2 px-4 transition duration-200 hover:bg-gray-700 hover:text-white"
			>
				asd
			</a>
			<a
				href="#"
				class="block py-2 px-4 transition duration-200 hover:bg-gray-700 hover:text-white"
			>
				asd
			</a>
			<a
				href="#"
				class="block py-2 px-4 transition duration-200 hover:bg-gray-700 hover:text-white"
			>
				asd
			</a>
		</nav> -->
	</aside>

	<main id="content" class="flex-1 p-6 lg:px-8">
		<div class="max-w-7xl mx-auto">
			<slot />
		</div>
	</main>
</div>
