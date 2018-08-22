var vm = new Vue({
    el: "#app",
    data: {
        title: "Hello Vue!"
    },
    filter: {

    },
    mounted: function () {
        this.$nextTick(function () {
            this.cartView();
        })
    },
    methods: {
        cartView: function () {
            this.title = "Molock.cn" 
        }
    }
})