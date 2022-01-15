
<Sidebar>
  <span slot="sidebar">
    <a href="/"     class="list-group-item list-group-item-action bg-light" use:link use:active>Home</a>
    <a href="/page" class="list-group-item list-group-item-action bg-light" use:link use:active>Example Page</a>
  </span>
  <span slot="content">
    <Router {routes}/>
  </span>
</Sidebar>

<script>
  import Router from 'svelte-spa-router'
  import {link, push, pop, replace, location, querystring} from 'svelte-spa-router'
  import active from 'svelte-spa-router/active'
  import Sidebar from './Sidebar.svelte'

  import Home from './Home.svelte'
  import Page from './Page.svelte'
  import NotFound from './NotFound.svelte'

  // routes
  const routes = {
    '/': Home,
    '/page' : Page,
    '*': NotFound,
  };

</script>

<style>
/* Style for "active" links; need to mark this :global because the router adds the class directly */
:global(.list-group-item.active) {
    font-weight: bold;
    color: black;
}


</style>
