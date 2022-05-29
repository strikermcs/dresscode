<template>
    <v-navigation-drawer v-model="drawer" color="#323B50" rail rail-width="90">
        <v-img :src="Logo" width="60" class="d-block text-center mx-auto mt-4"/>
        <v-list flat class="mt-5" style="background: #323B50" v-model="selectedItem">
            <v-list-item v-for="item in items" :key="item.id" 
            class="mb-8 item"
            @click="selectedItem = item.id">
                <div class="a" v-if="selectedItem == item.id"></div>
                <v-list-item-header>
                    <v-list-item-title class="text-center">
                        <v-icon :icon="item.icon" :color="selectedItem == item.id ? '#F34F64' : '#778290'"></v-icon>
                    </v-list-item-title>
                </v-list-item-header>
            </v-list-item>
        </v-list>
    </v-navigation-drawer>
</template>

<script>
import Logo from '../assets/logo.png'

export default {
    data: () => ({
        selectedItem: 0,
        drawer: null,
        Logo,
        items: [
            {id: 0, icon: 'mdi-laptop'},
            {id: 1, icon: 'mdi-tag-outline'},
            {id: 2, icon: 'mdi-shopping-outline'},
            {id: 3, icon: 'mdi-wallet-travel'},
            {id: 4, icon: 'mdi-apps'},
            {id: 5, icon: 'mdi-text-box-multiple-outline'}
        ]
    })
}
</script>


<style>

.a {
    position: absolute;
    top: 50%;
    left: 2%;
    transform: translate(-50%, -50%);
    height: 18px;
    width: 12px;
    border-radius: 0px 150px 150px 0px;
    background-color: #F34F64;
}

.item {
    cursor: pointer;
    transition: all 0.4s ease; 
}

</style>