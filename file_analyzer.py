import os

def count_characters_in_file(file_path):
    try:
        with open(file_path, 'r', encoding='utf-8') as file:
            content = file.read()
            return len(content)
    except Exception as e:
        print(f"Hata: {file_path} dosyası okunamadı - {str(e)}")
        return 0

def analyze_directory(directory):
    # İzlenecek dosya uzantıları
    target_extensions = {'.go', '.txt', '.mod', '.md'}
    
    # Sonuçları tutacak sözlük
    results = {ext: {'files': [], 'total_chars': 0} for ext in target_extensions}
    
    # Tüm dizini ve alt dizinleri gez
    for root, _, files in os.walk(directory):
        for file in files:
            _, ext = os.path.splitext(file)
            if ext in target_extensions:
                file_path = os.path.join(root, file)
                char_count = count_characters_in_file(file_path)
                results[ext]['files'].append(file_path)
                results[ext]['total_chars'] += char_count

    # Sonuçları yazdır
    print("\nDosya Analiz Sonuçları:")
    print("=" * 50)
    
    grand_total = 0
    for ext in target_extensions:
        file_count = len(results[ext]['files'])
        char_count = results[ext]['total_chars']
        grand_total += char_count
        
        print(f"\n{ext[1:].upper()} Dosyaları:")
        print(f"Dosya Sayısı: {file_count}")
        print(f"Toplam Karakter: {char_count:,}")
        
        # Dosya listesi
        if file_count > 0:
            print("\nDosyalar:")
            for file_path in results[ext]['files']:
                print(f"- {os.path.relpath(file_path, directory)}")
    
    print("\n" + "=" * 50)
    print(f"GENEL TOPLAM: {grand_total:,} karakter")

if __name__ == "__main__":
    # Şu anki dizini al
    current_dir = os.getcwd()
    print(f"Analiz edilen dizin: {current_dir}")
    analyze_directory(current_dir)