var path = require("path")
var webpack = require('webpack')
var BundleTracker = require('webpack-bundle-tracker')

module.exports = {
  context: __dirname,
  devtool: 'cheap-module-eval-source-map',

  entry: [
    'webpack-dev-server/client?http://localhost:3000',
    'webpack/hot/only-dev-server',
    './assets/js/index',
  ],

  output: {
    path: path.resolve('./assets/bundles/'),
    filename: "[name]-[hash].js",
    publicPath: "http://localhost:3000/assets/bundles/",
  },

  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    new webpack.NoErrorsPlugin(),
    new BundleTracker({filename: './webpack-stats.json'}),
  ],

  module: {
    loaders: [
      {
          test: /\.(js|jsx)$/,
          exclude: /node_modules/,
          loaders: ["react-hot", "babel-loader"],
      },{
          test: "/\.css$/",
          loaders: ['css'],
      }
    ],
  },

  resolve: {
    modulesDirectories: ['node_modules', 'bower_components'],
    extensions: ['', '.js', '.jsx']
  },
}
