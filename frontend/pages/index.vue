<template>
  <div class="app">
    <section class="todo">
      <header class="header">
        <h1>Task app</h1>
        <input class="new-todo" placeholder="put in your tasks" autofocus v-model="todo" @keyup.enter="addTodo">
      </header>
    </section>

    <div class="tasks-list">
      <section class="todo-list">
        <h2> Todo </h2>
        <section class="order-list">
          <div v-for="t in todos" :key="t.Id" class="item">
            <div v-if="t.Location === 'todo'">
              {{t.Text}}
            </div>
          </div>
        </section>
      </section>

      <section class="progress-list">
        <h2> Progress </h2>
        <section class="order-list">
          <div v-for="t in todos" :key="t.Id" class="item">
            <div v-if="t.Location === 'progress'">
              {{t.Text}}
            </div>
          </div>
        </section>
      </section>

      <section class="done-list">
        <h2> Done </h2>
        <section class="order-list">
          <div v-for="t in todos" :key="t.Id" class="item">
            <div v-if="t.Location === 'done'">
              {{t.text}}
            </div>
          </div>
        </section>
      </section>
    </div>
  </div>
</template>

<script>
import { mapMutations } from 'vuex';

export default {
  data () {
    return {
      todo: '',
    }
  },
  computed: {
    todos () { return this.$store.state.todos.todos }
  },
  methods: {
    addTodo () {
      this.$store.dispatch('todos/addTodo', this.todo)
      this.todo = ''
    },
    ...mapMutations({
      update: 'todos/update'
    })
  },
  mounted () {
    this.$store.dispatch('todos/getTodos')
  },

}

</script>
