<script>
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import Icon from '@iconify/svelte';

  const routes = [
    { path: '/dashboard', icon: 'lucide:layout-panel-left', label: 'Home' },
    { path: '/dashboard/storage', icon: 'lucide:database', label: 'storage' },
    { path: '/dashboard/logs', icon: 'lucide:line-chart', label: 'logs' },
    { path: '/dashboard/settings', icon: 'lucide:settings', label: 'Settings' },
  ];

  $: currentPath = $page.url.pathname;
</script>

<div class="main-layout">
<aside class="sidebar">
  <div class="nav-items">
    {#each routes as { path, icon, label }}
      <button
        class="nav-item"
        class:active={currentPath === path}
        aria-label={label}
        aria-current={currentPath === path ? 'page' : undefined}
        on:click={() => goto(path)}>
        <Icon icon={icon} height="20" width="20" />
      </button>
    {/each}
  </div>
</aside>

<slot />
</div>