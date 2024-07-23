module.exports = {
    // Diğer ayarlar...
    module: {
      rules: [
        // Diğer kurallar...
        {
          test: /\.js$/,
          enforce: 'pre',
          use: ['source-map-loader'],
          exclude: /node_modules\/stylis-plugin-rtl/
        },
      ],
    },
  };
  