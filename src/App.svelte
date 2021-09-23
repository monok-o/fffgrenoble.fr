<script>
  import Link from "./components/link.svelte";
  import Social from "./components/social.svelte";
  
  import Wave from "./components/wave.svelte";

  import linkFile from "./config/links.json"
  import socialFile from './config/socials.json'
</script>

<main>
  <Wave side="topleft" />
  <Wave side="bottomright" />

  <div class="content">
    <img class="part logo" src="/assets/logo.svg" alt="Fridays For Future Grenoble" />
    <div class="part links">
      <ul>
      {#each linkFile.links as link}
        <li><Link content="{link.content.toLocaleUpperCase()}" link="{link.link}" /></li>
      {/each}
      </ul>
    </div>
    <div class="part socials">
      {#each socialFile.socials as social}
        <Social icon="{social.icon}" link="{social.link}" />
      {/each}
    </div>
  </div>
</main>

<style>
	main {
		position: absolute;
		left: 0; right: 0;
		top: 0; bottom: 0;

    display: flex;
    align-items: center;
    justify-content: center;
	}

  .logo {
    max-width: 350px;
  }

  .content {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
  }

  .content .part {
    margin-top: 15px;
    margin-bottom: 15px;
  }

  ul {
    list-style: none;
    padding: 0;
  }

  @media only screen and (max-width: 600px) {
    .logo {
      max-width: 250px;
    }

    .content {
      margin: 15px;
    }
  }
</style>