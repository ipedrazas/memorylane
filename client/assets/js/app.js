

var app = new Vue({
    el: '#app',
    data: {
        message: "Hello Vue!"
    }
})

var apiURL = '/api/'
// var apiURL = 'http://192.168.64.29:31138/api/'
var listOfThings = new Vue({
    el: '#mylistofthings',
    data: {
        mythings: null
    },
    created: function () {
        this.fetchData();
    },
    methods: {
    fetchData: function () {
        var self = this;
        $.get( apiURL, function( data ) {
            self.mythings = data.results;
            });

        }
    }
})

// var eventsURL = 'http://192.168.64.29:31138/api/events/'
var eventsURL = '/api/events/'
var listOfEvents = new Vue({
    el: '#events',
    data: {
        events: null
    },
    created: function () {
        this.fetchData();
    },
    methods: {
        fetchData: function () {
            var self = this;
            $.get( eventsURL, function( data ) {
                self.events = data.results;
            });

        }
    }
})


// $.ajax({
//     url: 'http://127.0.0.1:5000/api/saveannotation',
//     type: 'POST',
//     contentType: 'application/json',
//     data: JSON.stringify({'sess_id' : $('#sessionid_area').val(),
//         'annotations': JSON.parse(annotations)}),
//     success: function(data) { alert(data.status); }
// });