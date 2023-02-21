<script lang="ts">
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import api from '$lib/api';
  import Menu from '$lib/components/menu.svelte';
  import notification from '$lib/components/notification/notification';
  import { onMount } from 'svelte';
  import { get } from 'svelte/store';
  import MenuAdd from './menu-add.svelte';

  let menu: any[] = [];
  let startAddMenu = false;

  page.subscribe((_) => {
    selectMenu();
  });

  function selectMenu() {
    let hash = get(page).url.hash;
    if (hash == '') return;

    let params = new URLSearchParams(hash.substring(1));
    const id = params.get('id');
    if (id == null) return;

    menu.forEach((el, i) => {
      menu[i].selected = el.id == +id;
      if (el.children != null) {
        el.children.forEach((child: any, j: any) => {
          menu[i].children[j].selected = child.id == +id;
        });
      }
    });
  }

  function moveMenu(id1: number, id2: number) {
    let seq1 = menu[id1].sequence,
      seq2 = menu[id2].sequence;

    api
      .post('admin/menu/reorder', [
        { id: menu[id1].id, sequence: seq2 },
        { id: menu[id2].id, sequence: seq1 },
      ])
      .then((resp) => {
        if (resp.status == 0) {
          menu[id1].sequence = seq2;
          menu[id2].sequence = seq1;
          menu.sort((a, b) => {
            return a.sequence - b.sequence;
          });
        } else notification.error(resp.message);
      })
      .catch((e) => notification.error(e));
  }

  function saveMenu(
    parentID: number,
    ev: CustomEvent,
    callback?: (success: boolean) => void
  ) {
    let name = ev.detail;

    api
      .postData('admin/menu/add', { parent_id: parentID, name: name })
      .then((data) => {
        goto('/admin/menu#id=' + data.id);
        refresh();
        if (callback != null) callback(true);
      })
      .catch((e) => {
        notification.error(e);
        if (callback != null) callback(false);
      })
      .finally(() => {
        if (parentID == 0) startAddMenu = false;
      });
  }

  export function refresh() {
    api.get('admin/menu').then((resp) => {
      if (resp.status == 0) {
        menu = resp.data;
        selectMenu();
      }
    });
  }

  onMount(() => refresh());
</script>

<div class="menu-list">
  {#each menu as m, i}
    <div>
      <Menu
        name={m.name}
        icon={m.icon}
        url="#id={m.id}"
        selected={m.selected}
      />
      {#if m.status == 'draft'}
        <i data-tooltip="draft" class="badge material-icons">edit</i>
      {:else if m.status == 'suspended'}
        <i data-tooltip="suspended" class="badge material-icons">pause</i>
      {/if}

      {#each m.children as sub}
        <div class="relative">
          <Menu
            name={sub.name}
            icon={sub.icon}
            url="#id={sub.id}"
            submenu={true}
            selected={sub.selected}
          />
          {#if sub.status == 'draft'}
            <i data-tooltip="draft" class="badge material-icons">edit</i>
          {:else if sub.status == 'suspended'}
            <i data-tooltip="suspended" class="badge material-icons">pause</i>
          {/if}
        </div>
      {/each}

      <div class="ml-4">
        <MenuAdd
          parentID={m.id}
          on:success={() => refresh()}
          tooltip="Add submenu"
        />
      </div>

      {#if m.sequence > 0}
        <div class="move-up" data-tooltip="Move up">
          <button class="material-icons" on:click={() => moveMenu(i - 1, i)}>
            arrow_right
          </button>
        </div>
      {/if}
      {#if i != menu.length - 1}
        <div class="move-down" data-tooltip="Move down">
          <button class="material-icons" on:click={() => moveMenu(i, i + 1)}>
            arrow_left
          </button>
        </div>
      {/if}
    </div>
  {/each}

  <div>
    <MenuAdd parentID={0} on:success={() => refresh()} tooltip="Add menu" />
  </div>
</div>

<style lang="scss">
  @import '../../../../styles/colors';

  .menu-list {
    @apply w-fit space-y-1;

    & > div {
      @apply relative border-2 rounded-lg flex-col p-1 items-center pl-8;
      min-width: 18em;
    }

    & .badge {
      @apply absolute tooltip-right cursor-default top-3 right-2 text-red-500 text-sm;
    }
  }

  .btn-sequence {
    background-color: transparent;
    color: #666;
    box-shadow: 0 0 4px -2px #000;
    &:hover {
      box-shadow: 0 0 4px 0px $color-secondary;
      background-color: $color-secondary;
      color: $color-text;
    }
  }

  .move-up,
  .move-down {
    @apply absolute w-fit left-2;
    & button {
      @apply p-0 transition-colors rounded-full -rotate-90;
      background-color: transparent;
      color: $color-primary;
    }
  }

  .move-up {
    @apply top-2;
  }

  .move-down {
    @apply bottom-2;
  }

  .add-submenu {
    margin: 0.5rem 0.5rem 0.5rem 2rem;
  }
</style>
