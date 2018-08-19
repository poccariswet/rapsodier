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
          <div v-for="t in todos" :key="t.id" class="item">
            <div v-if="t.location === 'todo'">
              {{t.text}}
            </div>
          </div>
        </section>
      </section>

      <section class="progress-list">
        <h2> Progress </h2>
        <section class="order-list">
          <div v-for="t in todos" :key="t.id" class="item">
            <div v-if="t.location === 'progress'">
              {{t.text}}
            </div>
          </div>
        </section>
      </section>

      <section class="done-list">
        <h2> Done </h2>
        <section class="order-list">
          <div v-for="t in todos" :key="t.id" class="item">
            <div v-if="t.location === 'done'">
              {{t.text}}
            </div>
          </div>
        </section>
      </section>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      todo: '',
      todos: [
        {
          id: 0,
          text: "sample text",
          done: false,
          location: "todo",
          created_at: "2018-01-01"
        }
      ]
    }
  },
  methods: {
    addTodo () {
//      const obj = {
//        text: this.todo
//      }
//      const method = 'POST'
//      const body = JSON.stringify(obj)
//      const headers = {
//                'Content-Type': 'application/json',
//            };
//      fetch("localhost:9999/tasks", {
//          method,
//          headers,
//          body
//      }).then((res) => res.json()).then((res) => {
//        console.log(res)
//      }).catch((err) => {
//        console.log(err)
//      });

      this.todo = ''
    }
  }
}

</script>
