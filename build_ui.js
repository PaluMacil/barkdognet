const esbuild = require('esbuild');
const sassPlugin = require('esbuild-sass-plugin').default;

const isWatchMode = process.argv.includes('--watch');
const buildOptions = {
    entryPoints: ['src/index.ts'],
    bundle: true,
    outdir: 'dist/ui/',
    sourcemap: true,
    splitting: true,
    format: 'esm',
    minify: true,
    plugins: [sassPlugin()],
};

if (isWatchMode) {
    void esbuild.context(buildOptions)
        .then(async (ctx) => {
            await ctx.watch();
        }).catch((err) => {
            console.error('Error occurred during watch', err);
        })
} else {
    esbuild.build(buildOptions).catch(() => process.exit(1));
}

