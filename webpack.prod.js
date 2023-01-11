import { merge } from 'webpack-merge';
import MiniCssExtractPlugin from 'mini-css-extract-plugin';
import CssMinimizerPlugin from 'css-minimizer-webpack-plugin';
import common from './webpack.common.js';

const [mainCommonConfig, swConfig] = common;

const mainConfig = merge(mainCommonConfig, {
  mode: 'production',
  output: {
    filename: '[name].[contenthash].mjs',
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: '[name].[contenthash].css',
    }),
    new CssMinimizerPlugin()
  ],
});

export default [mainConfig, swConfig];