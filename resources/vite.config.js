export default {
    base: '/assets/dist',
    build: {
        outDir: '../assets/dist',
        rollupOptions: {
            input: 'src/index.js',
            output: {
                entryFileNames: 'index.js',
                assetFileNames: '[name].[ext]',
            }
        },
    },
    plugins: [],
}