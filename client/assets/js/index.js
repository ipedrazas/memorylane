// import Vue from 'vue'
// import App from './App'

// import fullCalendar from 'vue-fullcalendar'

// Vue.component('full-calendar', fullCalendar)

// Vue2
// new Vue({
//   el : '#app',
//   render: h => h(App),
//   template : '<App/>',
//   components : {
//     App
//   }
// })

//Vue1

// new Vue({
// 	el : 'body',
// 	components : {
// 		App
// 	}
// })

var demoEvents = [
	{
      title : 'Sunny Out of Office',
      start : '2017-05-5',
      end : '2017-05-17'
    }
]

// $(function () {
    Vue.component('full-calendar', VueFullcalendar)

    var app = new Vue({
        el: '#app',
        data () {
            return {
            fcEvents : demoEvents
            }
        }
    })
// })
