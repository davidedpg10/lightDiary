  <div class="d-flex" id="wrapper" class:toggled={menuActive}>
    <div class="bg-light border-right" id="sidebar-wrapper">
      <slot name="sidebar-heading">
	<div class="sidebar-heading">
	  Welcome!
	</div>
      </slot>
      <div class="list-group list-group-flush">
	<slot name="sidebar">
	</slot>
      </div>

    </div>
    <div id="page-content-wrapper">
      <nav class="navbar  navbar-light bg-light border-bottom">

	<button class="navbar-toggler" id="menu-toggle" on:click={() => menuActive=!menuActive}>
	  <span class="navbar-toggler-icon"></span>
	</button>
	<slot name="top-content">
	  <div class="navbar-brand"> Hello </div>
	</slot>
      </nav>
      <div class="container-fluid">
	<slot name="content">
      	</slot>

      </div>
    </div>
  </div>

<script>
  let menuActive=false;
</script>
