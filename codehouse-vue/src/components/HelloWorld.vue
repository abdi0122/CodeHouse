<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h3>Nav Links</h3>
    <ul>
      <li>
        <router-link to=""></router-link>
      </li>
    </ul>
    <p>
      Women are often times excluded in conversations relating to technology,
      and often times there are very few women in these spaces, especially
      classrooms, to begin with. This website serves as a hub to provide
      students with a place to come together and be heard, as well as find other
      women in stem to study with based on similar coursework.
    </p>
    <button>Login</button>
    <button>Sign Up</button>
  </div>
</template>

<script>
export default {
  name: "HelloWorld",
  props: {
    msg: String,
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
