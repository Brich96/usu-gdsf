<template>
    <v-container data-test="container" :style="containerStyle">
        <v-row
            class="fill-height"
            align-content="center"
            justify="center"
        >
            <v-col
                data-test="text-wrapper"
                class="text-subtitle-1 text-center"
                cols="12"
            >
                Loading {{ text }}
            </v-col>
            <v-col cols="6" data-test="progress-wrapper">
                <v-progress-linear
                    color="primary"
                    indeterminate
                    rounded
                    height="6"
                ></v-progress-linear>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
    export default {
        name: 'Loading',
        props: {
            text:String,
            containerStyle:String
        }
    }
</script>