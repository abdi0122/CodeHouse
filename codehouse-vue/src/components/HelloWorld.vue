<template>
  <div class="hello">
    <h1>Study 2gether</h1>
    <p>
      Women are often times excluded in conversations relating to technology,
      and often times there are very few women in these spaces especially
      classrooms, to begin with. This website will provide students with a place
      to come together and be heard, as well as find other women in stem to
      study with based on similar coursework.
    </p>
    <router-link to="/about"><button>Login</button></router-link> |
    <router-link to="/about"><button>Sign Up</button></router-link>
  </div>
</template>

<script>
export default {
  name: "HelloWorld",
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
button {
  background-color: #3A405A;
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;
}
</style>
