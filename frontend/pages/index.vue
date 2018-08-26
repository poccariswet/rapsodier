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
          <draggable element="ul" :options="{ group: 'tasks' }">
            <div v-for="t in tasks.todos" :key="t.Id" class="item">
              <li v-if="t.Location === 'todo'">{{t.Text}}</li>
            </div>
          </draggable>
        </section>
      </section>

      <section class="progress-list">
        <h2> Progress </h2>
        <section class="order-list">
          <draggable element="ul" :options="{ group: 'tasks' }">
            <div v-for="t in tasks.progresses" :key="t.Id" class="item">
              <li v-if="t.Location === 'progress'">{{t.Text}}</li>
            </div>
          </draggable>
        </section>
      </section>

      <section class="done-list">
        <h2> Done </h2>
        <section class="order-list">
          <draggable element="ul" :options="{ group: 'tasks' }">
            <div v-for="t in tasks.dones" :key="t.Id" class="item">
              <li v-if="t.Location === 'done'">{{t.Text}}</li>
            </div>
          </draggable>
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
    tasks () { return this.$store.state.todos.tasks }
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
