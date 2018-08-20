export const state = () => ({
  todos: []
})

export const mutations = {

  setTodos(state, todos) {
    state.todos = todos
  },

  add(state, text) {
    state.todos.push({
      id: state.list.length + 1,
      text: text,
      done: false,
      location: 'todo',
      created_at: Date.now()
    })
  },

  update(state, todo) {
    todo.done = !todo.done
  }

}

export const actions = {
  getTodos({commit}) {
    const method = 'GET'
    const headers = {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
    }

    fetch('http://localhost:9999/tasks',{
      method,
      headers
    }).then((res) => res.json()).then((res) => {
        commit("setTodos", res)
    }).catch(console.error)
  },

  addTodo({commit}, obj) {
    const method = 'POST'
    const body = JSON.stringify(obj)
    const headers = {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
    }

    fetch("http://localhost:9999/tasks", {
        method,
        headers,
        body
    }).then((res) => res.json()).then((res) => {
      console.log(res)
      commit('add', obj.text)
    }).catch((err) => {
      console.log(err)
    })
  }
}
