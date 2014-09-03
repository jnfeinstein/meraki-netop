module.exports = function(grunt) {

  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    uglify: {
      options: {
        sourceMap: true
      },
      dist: {
        files: {
          'public/javascripts/build.min.js': 'private/javascripts/**/*.js'
        }
      }
    },
    compass: {
      dist: {
        options: {
          sassDir: 'private/stylesheets',
          cssDir: 'public/stylesheets'
        }
      }
    },
    watch: {
      stylesheets: {
        files: ['private/stylesheets/**/*.scss'],
        tasks: ['compass']
      },
      javascripts: {
        files: ['private/javascripts/**/*.js'],
        tasks: ['uglify']
      }
    },
    clean: [
      'public/javascripts/**/*.min.js',
      'public/javascripts/**/*.map',
      'public/stylesheets/**/*.css',
      'public/stylesheets/**/*.css.map',
    ]
  });

  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-compass');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-clean');

  grunt.registerTask('default', ['clean', 'uglify', 'compass']);
  grunt.registerTask('heroku', ['clean', 'uglify', 'compass']);

};