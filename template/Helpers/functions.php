<?php
use Illuminate\Support\HtmlString;

if (! function_exists('asset_hash')) {
    /**
     * Get the path to a versioned Mix file.
     *
     * @param  string  $path
     * @param  string  $manifestDirectory
     * @param  string  $manifestFileName
     * @return \Illuminate\Support\HtmlString
     *
     * @throws \Exception
     */
    function asset_hash($path, $manifestDirectory = 'static', $manifestFileName = 'manifest.json')
    {
        static $manifests = [];
        
        if (! starts_with($path, '/')) {
            $path = "/{$path}";
        }
        
        if ($manifestDirectory && ! starts_with($manifestDirectory, '/')) {
            $manifestDirectory = "/{$manifestDirectory}";
        }

//        if (file_exists(public_path($manifestDirectory.'/hot'))) {
//            return new HtmlString("//localhost:8080{$path}");
//        }
        
        $manifestPath = public_path($manifestDirectory.'/'.$manifestFileName);
        
        if (! isset($manifests[$manifestPath])) {
            if (! file_exists($manifestPath)) {
                // throw new Exception('The Mix manifest does not exist.');
                return new HtmlString($path);
            }
            
            $manifests[$manifestPath] = json_decode(file_get_contents($manifestPath), true);
        }
        
        $manifest = $manifests[$manifestPath];
        
        if (! isset($manifest[$path])) {
            throw new Exception(
                "Unable to locate Mix file: {$path}. Please check your ".
                'webpack.mix.js output paths and try again.'
            );
        }
        
        // return new HtmlString($manifestDirectory.$manifest[$path]);
        return new HtmlString($manifest[$path]);
    }
}
