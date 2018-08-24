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
            <ul v-if="t.Location === 'todo'">
              <draggable :options="{ group: 'todos' }">
                <li>{{t.Text}}</li>
              </draggable>
            </ul>
          </div>
        </section>
      </section>

      <section class="progress-list">
        <h2> Progress </h2>
        <section class="order-list">
          <div v-for="t in todos" :key="t.Id" class="item">
            <ul v-if="t.Location === 'progress'">
              <draggable :options="{ group: 'todos' }">
                <li>{{t.Text}}</li>
              </draggable>
            </ul>
          </div>
        </section>
      </section>

      <section class="done-list">
        <h2> Done </h2>
        <section class="order-list">
          <div v-for="t in todos" :key="t.Id" class="item">
            <ul v-if="t.Location === 'done'">
              <draggable :options="{ group: 'todos' }">
                <li>{{t.Text}}</li>
              </draggable>
            </ul>
          </div>
        </section>
      </section>
    </div>
  </div>
</template>

<script>
import { mapMutations } from 'vuex'
import draggable from 'vuedraggable'

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
  components: {
    draggable,
  }

}

</script>
