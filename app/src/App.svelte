<script>
  import Counter from './lib/Counter.svelte'
  import Login from "./lib/Login.svelte"
  import { Router, Link, Route } from "svelte-routing";
  import Home from "./routes/Home.svelte";
  import About from "./routes/About.svelte";
  import Blog from "./routes/Blog.svelte";
  import BlogPost from "./routes/BlogPost.svelte"
  import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query'

  const queryClient = new QueryClient()
  const url = ""
</script>

<main>
  <QueryClientProvider client={queryClient}>
    <Router url="{url}">
      <nav>
        <Link to="/">Home</Link>
        <Link to="about">About</Link>
        <Link to="blog">Blog</Link>
      </nav>
      <div>
        <Route path="blog/:id" component="{BlogPost}" />
        <Route path="blog" component="{Blog}" />
        <Route path="about" component="{About}" />
        <Route path="/"><Home /></Route>
      </div>
    </Router>
  </QueryClientProvider>

</main>

<style>
</style>
