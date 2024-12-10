$files = Get-ChildItem -Path "functions" -Recurse -Filter "*.go"
foreach ($file in $files) {
    $content = Get-Content $file.FullName -Raw
    $newContent = $content -replace 'github.com/kha0sys/nodo.social/(?!functions)', 'github.com/kha0sys/nodo.social/functions/'
    Set-Content -Path $file.FullName -Value $newContent
}
