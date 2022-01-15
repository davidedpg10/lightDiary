<script>
    import FaHome from 'svelte-icons/fa/FaHome.svelte'
    import FaGithub from 'svelte-icons/fa/FaGithub.svelte'
</script>

<nav>
	<a href="/" class="icon"><FaHome />Home</a>
	<a href="https://github.com/davidedpg10/golang-diary" class="icon"><FaGithub />Github</a>
</nav>

<style>
    nav {
        background: rgb(216, 243, 255);
        padding: 1em;
    }

    nav a {
        color: black;
        font-weight: bold;
        display: inline-block;
        padding: 1;
        margin-left: 3em;
        margin-right: 3em;
        text-decoration: none;
        font-size: large;
        justify-content: center;
    }
    .icon {
    width: 32px;
    height: 32px;
  }
</style>

<slot></slot>