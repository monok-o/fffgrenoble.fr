<script>
  import Link from "./components/link.svelte";
  import Social from "./components/social.svelte";
  import Event from "./components/event.svelte"

  import Wave from "./components/wave.svelte";

  import linkFile from "./config/links.json"
  import socialFile from './config/socials.json'
  import eventFile from './config/events.json'

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

    <div class="part events">
      <h2 class="events title">Événements à venir :</h2>
      <hr class="first_separator">
      <ul>
          {#each eventFile.events as event}
            <li class="event"><Event title={event.title} date={event.date} time={event.time} place={event.place} link={event.link}/></li>
          {/each}
      </ul>
    </div>

  </div>
</main>

<style>
	main {
		position: absolute;
		left: 0; right: 0;
		top: 0; bottom: 0;

    display: flex;
    justify-content: center;
	}

  .logo {
    max-width: 350px;
  }

  .content {
    position: absolute;
    top: 50px;

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

  .events{
    width:100%;
  }

  .events .title {
    color:white;
  }

  .events .first_separator {
    background-color:black;
    border : 1.5px solid white;
  }

  .event {
    border-bottom: 1.5px solid white;
  }

  @media only screen and (max-width: 600px) {
    .logo {
      max-width: 250px;
    }

    .content {
      margin: 15px;
    }
  }

  @media only screen and (max-height: 520px) {
    .part {
      margin-top: 0px;
      margin-bottom: 0px;
    }
  }

  
</style>