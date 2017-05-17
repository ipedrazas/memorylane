

var demoEvents = [
	{
      title : 'Camelot',
      start : '2017-05-17',
      end : '2017-05-19',
      cssClass  : 'camelot'
    },
    {
      title : 'SH',
      start : '2017-05-15',
      end : '2017-05-16',
      cssClass  : 'sohohouse'
    }
]


Vue.component('full-calendar', VueFullcalendar)

var app = new Vue({
    el: '#app',
    data () {
        return {
        fcEvents : demoEvents
        }
    }
})

