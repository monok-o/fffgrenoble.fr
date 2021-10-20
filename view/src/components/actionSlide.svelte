<script>
  import {FullpageSlide} from 'svelte-fullpage'

  export let title
  export let backgroundImage
  export let description
</script>

<FullpageSlide class='container' center>
  <img src="assets/images/{backgroundImage}" alt="stand">
  <div class="container-2">  
    <h2 class='title'>{title}</h2>
    <p class='description'>{description}</p>
  </div>
</FullpageSlide>

<style scoped>

.title {
  font-size: 30px;
  font-weight: 700;
  margin:10px 0;
  color:white;
  text-align: center;
}

.description{
  font-size: 20px;
  font-weight: 400;
  color:white;
  text-align: center;
}

img {
  position: absolute;
  top:0;
  left: 0;
  z-index: -2;
  filter:brightness(30%);
  width:100%;
  height:100%;
  object-fit: cover;

}

.container-2 {
  width:30%;
  flex-direction:column;
  justify-content: space-around;
  z-index: 0;
}

@media only screen and (max-width: 540px) {
  .title {
    font-size: 20px;
  }
  
}
@media only screen and (max-width: 800px) {
  
  .container-2 {
    width:60%
  }
}

@media only screen and (max-width: 330px) {
  .title {
    font-size: 15px;
  }
}

@media only screen and (max-width: 330px) {
  
  .container-2 {
    width:90%
  }
}


</style>