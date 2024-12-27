<script>
  import { onMount } from 'svelte';

  let storageUsed = 65;
  let storageTotal = 100;
  let uptime = 99.9;
  let recentlyAdded = [
    { name: 'idk placeholder', size: '2.3 MB', date: 'later' },
  ];
</script>

<div class="app-body">
  <section class="overview">
    <div class="storage box">
      <h1>Storage</h1>
      <div class="storage-bar">
        <div class="storage-used" style="width: {storageUsed}%"></div>
      </div>
      <div class="storage-text">
        {storageUsed}GB / {storageTotal}GB
      </div>
    </div>
    <div class="uptime box">
      <h1>Uptime</h1>
      <div class="uptime-value">{uptime}%</div>
    </div>
  </section>
  <section class="recents-box">
    <h1>Recently Added</h1>
    <div class="contents">
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Size</th>
            <th>Date</th>
          </tr>
        </thead>
        <tbody>
          {#each recentlyAdded as item}
            <tr>
              <td>{item.name}</td>
              <td>{item.size}</td>
              <td>{item.date}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </section>
</div>
