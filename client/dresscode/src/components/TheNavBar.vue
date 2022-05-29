<template>
    <v-app-bar app dark color="#323B50" flat>
        <v-tabs color="#F34F64">
            <v-tabs-slider color="#F34F64"></v-tabs-slider>
            <v-tab class="withoutupercase">Shirt</v-tab>
            <v-tab class="withoutupercase">Shoes</v-tab>
            <v-tab class="withoutupercase">Jackets</v-tab>
            <v-tab class="withoutupercase">Hats</v-tab>
            <v-tab class="withoutupercase">Trousers</v-tab>
        </v-tabs>
        <v-btn size="small" color="#F34F64" variant="contained">
            <v-icon icon="mdi-chevron-down"></v-icon>
        </v-btn>
        <v-spacer></v-spacer>
        <v-divider vertical inset class="mr-4" color="#778290"></v-divider>
        <v-text-field label="Search" 
        prepend-inner-icon="mdi-magnify" 
        dense class="mt-10"
        bg-color="#323B50"
        color="#778290"
        variant="plain"
        ></v-text-field>
    </v-app-bar>
</template>


<script>
export default {
    data: () => ({
        links: [
            'Dashboard',
            'Messages',
            'Profile',
            'Updates'
        ],
        items : [
            {icon: 'mdi-laptop'},
            {icon: 'mdi-tag-outline'},
            {icon: 'mdi-shopping-outline'},
            {icon: 'mdi-wallet-travel'},
            {icon: 'mdi-apps'},
            {icon: 'mdi-text-box-multiple-outline'},
        ]
    })
}
</script>

<style>
.v-tab__slider {
    position: absolute;
    bottom: -8px !important;
    left: 50% !important;
    transform: translate(-50%, -50%);
    height: 12px !important;
    width: 20px !important;
    border-radius: 150px 150px 0px 0px;
    border-color: #F34F64;
}

.v-tab.withoutupercase {
    text-transform: none !important;
    color: #778290;
}


.v-tabs {
    width: 30% !important;
}

.v-btn:not(.v-btn--round).v-size--small {
    min-width: 10px !important;
    padding: 0 5px !important;
}
</style>